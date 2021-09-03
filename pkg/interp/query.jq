# []
def _query_array:
  {
    term: {
      type: "TermTypeArray",
      array: {
        query: .
      }
    }
  };

# a() -> b()
def _query_func_rename(name):
  .term.func.name = name;

# . | r
def _query_pipe(r):
  { op: "|",
    left: .,
    right: r
  };

def _query_ident: {term: {type: "TermTypeIdentity"}};

# .[]
def _query_iter:
  { "term": {
      "suffix_list": [{
        "iter": true
      }],
      "type": "TermTypeIdentity"
    }
  };

def _query_is_func(name):
  .term.func.name == name;

# last query in pipeline
def _query_pipe_last:
  if .term.suffix_list then
    ( .term.suffix_list[-1]
    | if .bind.body then
        ( .bind.body
        | _query_pipe_last
        )
      end
    )
  elif .op == "|" then
    ( .right
    | _query_pipe_last
    )
  end;

def _query_transform_last(f):
  def _f:
    if .term.suffix_list then
      .term.suffix_list[-1] |=
        if .bind.body then
          .bind.body |= _f
        else f
        end
    elif .op == "|" then
      .right |= _f
    else f
    end;
  _f;

# <filter...> | <slurp_func> -> [.[] | <filter...> | .] | (<slurp_func> | f)
def _query_slurp_wrap(f):
  ( . as {$meta, $imports}
  # move directives new root query
  | del(.meta)
  | del(.imports)
  | _query_pipe_last as $lq
  | _query_transform_last(_query_ident) as $pipe
  | _query_iter
  | _query_pipe($pipe)
  | _query_array
  | _query_pipe($lq | f)
  | .meta = $meta
  | .imports = $imports
  );

#compdef npkg

# ~/.local/share/zsh/completions/_npkg
#
# modified version of cli's template:
#   https://github.com/urfave/cli/blob/main/autocomplete/zsh_autocomplete

local -a opts
local cur
cur=${words[-1]}
if [[ "$cur" == "-"* ]]; then
  opts=("${(@f)$(${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
else
  opts=("${(@f)$(${words[@]:0:#words[@]-1} --generate-bash-completion)}")
fi

if [[ "${opts[1]}" != "" ]]; then
  _describe 'values' opts
else
  _files
fi

# Model

machina installs the [model](https://github.com/krisyotam/model)
standard. It does not replace it and it does not copy the templates
into this repo.

model is how every agent CLI on a machine shares the same facts:
hosts, how to invoke the other CLIs, a slug index, a private env file,
and a session log. The contract is one line in each agent's instruction
file. The rest is Markdown in `$XDG_CONFIG_HOME/model/`.

## What install does

1. Find a checkout of `github.com/krisyotam/model`.
   - `$MACHINA_MODEL_REPO` if set
   - `$HOME/src/model` if that directory looks like the repo
   - otherwise clone into `$XDG_CACHE_HOME/machina/model`
2. Create `$XDG_CONFIG_HOME/model/` and `$XDG_STATE_HOME/model/`.
3. Copy from `template/` any file that is not already there.
   `models.env.example` becomes `models.env` at mode `0600`.
4. Copy `template/log.csv` into state if `log.csv` is missing.
5. Append the model one-liner to each agent instruction file that
   does not already mention the config path:

   ```
   Shared primitives (hosts/SSH, subagent CLIs, models.env) live in `~/.config/model/`. Read `~/.config/model/README.md` on launch and append a row to `~/.local/state/model/log.csv` when the session ends.
   ```

The paths in that line are the resolved ones (`~/.config/model` or
whatever `$XDG_CONFIG_HOME` expanded to). Agents do not expand
`$XDG_*` themselves.

## What install will not do

- Overwrite `hosts.md` or `models.env`. Those are local facts and
  secrets.
- Overwrite any other model file that already exists.
- Fill in `hosts.md` or `models.env`. A human does that.
- Commit `models.env` anywhere.

If the clone fails, `machina install model` prints the git URL and
stops. The rest of `machina install` can still run (`dirs`, `bin`).

## After install

Fill `$XDG_CONFIG_HOME/model/hosts.md` and `models.env`. The templates
in the model repo show the shape. Real values come from Infisical or
`pass`, never from this repo.

The session log is `$XDG_STATE_HOME/model/log.csv`. Every agent
appends one row when a session ends. That rule is model's, not
machina's; the one-liner is how the agent finds it.

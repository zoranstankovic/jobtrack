# 03: Edit and delete Companies inline

**What to build:** In the Companies list the user clicks Edit on a row. The row turns into an inline form, and saving swaps the updated row back in without a reload. The user can also delete a Company after confirming, and the row disappears. (The rule that a Company with Applications cannot be deleted comes in 06, once Applications exist.)

**Teaches:** HTMX `hx-get` / `hx-put` / `hx-delete`, swap targets and `outerHTML` swaps, edit and read modes for one row, `hx-confirm`.

**Blocked by:** 02

**Status:** ready-for-agent

- [ ] Clicking Edit replaces the row with an inline form prefilled with the Company's current values
- [ ] Saving updates the Company and swaps the read-only row back; Cancel restores the row unchanged
- [ ] Validation (required name, unique name ignoring case, excluding the Company itself) works the same as on create
- [ ] Delete asks for confirmation, then removes the Company and its row
- [ ] Tests through the HTTP seam cover: edit, edit validation error, delete

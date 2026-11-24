# 13: Dashboard

**What to build:** The Dashboard becomes the home page with two sections. "New postings" shows Job Postings with Posting Status New, newest first, with their Search Profiles, the Dismiss and Create Application buttons, and the "Refresh all" button. "Active applications" lists every Application not in Accepted, Rejected, Withdrawn or Ghosted, oldest change first, showing Company, position, current Application Status and the number of days since the last Status Change. Changing a status (with the inline form from 05) moves or removes the Application in the list without a reload. So do Dismiss and Refresh.

**Teaches:** composing a page from reusable fragments, HTMX events (`HX-Trigger` response headers and `hx-trigger` listeners) so that one action updates another section, date arithmetic in SQL or Go.

**Blocked by:** 05, 10

**Status:** ready-for-agent

- [ ] The root URL shows the Dashboard
- [ ] "New postings" reuses the New postings list from 10 with Refresh all, Dismiss and Create Application
- [ ] "Active applications" excludes Accepted, Rejected, Withdrawn and Ghosted, is sorted by last change (oldest first) and shows the days since the last Status Change
- [ ] Changing an Application's status from the Dashboard updates its row, or removes it when the status becomes final, without a reload
- [ ] Refresh all and Dismiss update the New postings section in place
- [ ] Tests through the HTTP seam cover: days since last change using Status Changes dated in the past, final statuses excluded, sort order

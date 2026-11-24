# 10: New postings list: Refresh all and Dismiss

**What to build:** A New postings view lists every Job Posting with Posting Status New, newest first, each showing which Search Profiles found it. "Refresh all" runs every Search Profile and updates the list in place, with a total of new postings. Dismiss on a posting removes it from the list immediately. A Dismissed Job Posting never comes back, even when a later Refresh finds it again. The Search Profiles list shows how many New postings each profile has.

**Teaches:** many-to-many queries (aggregating the profile names for each posting), HTMX element removal, running one operation over many items and collecting the results.

**Blocked by:** 09

**Status:** ready-for-agent

- [ ] The New postings view shows only Posting Status New, newest publication first, with the names of the Search Profiles that found each posting
- [ ] "Refresh all" refreshes every Search Profile and re-renders the list with the total number of new postings; a failing profile is reported without stopping the others
- [ ] Dismiss sets Posting Status to Dismissed and removes the element without a reload
- [ ] A Refresh whose results contain a Dismissed posting does not show it again
- [ ] The Search Profiles list shows the count of New postings per profile
- [ ] Tests through the HTTP seam (with the fake BA server) cover: Refresh all across two profiles sharing a posting, Dismiss, a Dismissed posting staying hidden after a Refresh

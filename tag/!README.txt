Github thinks the project's main language is C++ because
there are lots of .h files in the include/ directory.
Thus godoc.org cannot index this project, as it looks
only for projects tagged as Go. The solution is to create
99 .go files so that github knows this project's main
language is Go.

	var diff = `diff --git a/newfile2 b/newfile2
	var diff2 = `diff --git "a/A \\ B" "b/A \\ B"
	var diff2a = `diff --git "a/A \\ B" b/A/B
	var diff3 = `diff --git a/README.md b/README.md
	assert.NoError(t, diff.LoadComments(issue, user))
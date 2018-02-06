default:
	echo doing nothing

update-changelog:
	markdown-toc -min 2 -max 2 CHANGELOG.md

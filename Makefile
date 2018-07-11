default:
	echo doing nothing

update-docs:
	markdown-toc -min 2 -max 2 README.md
	markdown-toc -min 2 -max 2 CHANGELOG.md

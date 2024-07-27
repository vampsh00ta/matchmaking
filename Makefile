PROJECTNAME=matchmaking

init:
	find . -type  f -exec sed -i '' -e's/matchmaking/$(PROJECTNAME)/g' {} + -not -path "Makefile"

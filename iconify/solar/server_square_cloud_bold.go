package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerSquareCloudBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="solarServerSquareCloudBold0" fill-rule="evenodd" d="M20.654 16.88c-.111-1.057-.958-1.88-1.987-1.88c-.845 0-1.568.555-1.861 1.34a2.222 2.222 0 0 0-.015 1.513a1.46 1.46 0 0 0-1.4.49A1.633 1.633 0 0 0 15 19.412c0 .877.672 1.588 1.5 1.588H20c1.105 0 2-.948 2-2.118c0-.927-.563-1.714-1.346-2.001Z" clip-rule="evenodd"/></defs><g fill="currentColor"><path fill-rule="evenodd" d="M3.172 19.828C4.343 21 6.229 21 10 21h3.93a3.165 3.165 0 0 1-.43-1.588c0-.816.32-1.592.857-2.162H13.5a.75.75 0 0 1 0-1.5h1.928c.511-1.29 1.727-2.25 3.239-2.25c1.56 0 2.81 1.028 3.288 2.38l.014.01C22 15.073 22 14.12 22 13v-.25H2V13c0 3.771 0 5.657 1.172 6.828ZM6 18.25a.75.75 0 0 1-.75-.75v-2a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-.75.75Zm3 0a.75.75 0 0 1-.75-.75v-2a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-.75.75ZM3.172 4.172C2 5.343 2 7.229 2 11v.25h20V11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172ZM9 9.25a.75.75 0 0 1-.75-.75v-2a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-.75.75ZM5.25 8.5a.75.75 0 0 0 1.5 0v-2a.75.75 0 0 0-1.5 0v2Zm7.5-1a.75.75 0 0 1 .75-.75H18a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/><use href="#solarServerSquareCloudBold0" fill-rule="evenodd" clip-rule="evenodd"/><use href="#solarServerSquareCloudBold0"/></g>`),
		g.Group(children),
	)
}
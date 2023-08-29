package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasswordMinimalisticInputBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.75 2a.75.75 0 0 0-1.5 0v20a.75.75 0 0 0 1.5 0v-2.006c2.636-.027 4.104-.191 5.078-1.166C22 17.657 22 15.771 22 12c0-3.771 0-5.657-1.172-6.828c-.974-.975-2.442-1.139-5.078-1.166V2Z"/><path fill-rule="evenodd" d="M3.172 18.828C4.343 20 6.229 20 10 20h3V4h-3C6.229 4 4.343 4 3.172 5.172C2 6.343 2 8.229 2 12c0 3.771 0 5.657 1.172 6.828ZM13 12a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm-4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
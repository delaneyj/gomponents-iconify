package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17.334V6.667c0-.88 0-1.32.185-1.58a1 1 0 0 1 .687-.412c.317-.04.705.166 1.48.58l10 5.333l.004.002c.857.457 1.286.686 1.427.99c.122.266.122.573 0 .839c-.141.305-.571.535-1.43.993l-10 5.333c-.777.414-1.164.62-1.48.58a1 1 0 0 1-.688-.412C5 18.653 5 18.213 5 17.333Z"/>`),
		g.Group(children),
	)
}
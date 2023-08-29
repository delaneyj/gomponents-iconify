package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orionviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.093 38.464a9.99 9.99 0 1 0-5.052-19.279"/><circle cx="26.715" cy="7.864" r=".75" fill="currentColor"/><circle cx="22.676" cy="25.949" r=".75" fill="currentColor"/><circle cx="35.576" cy="23.25" r=".75" fill="currentColor"/><circle cx="40.742" cy="28.766" r=".75" fill="currentColor"/><circle cx="19.764" cy="41.089" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.24 31.14a5.307 5.307 0 0 0 5.307 5.307m0 0a5.307 5.307 0 0 0-5.307 5.307"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.24 41.755a5.307 5.307 0 0 0-5.307-5.308"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.932 36.447a5.307 5.307 0 0 0 5.308-5.307m8.919-18.328a3.333 3.333 0 0 0 3.333 3.333m0 0a3.333 3.333 0 0 0-3.333 3.333m0 0a3.333 3.333 0 0 0-3.333-3.333m0 0a3.333 3.333 0 0 0 3.333-3.333M15.972 7.487a4.542 4.542 0 0 0 4.542 4.542m0 .001a4.542 4.542 0 0 0-4.542 4.542"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.972 16.572a4.542 4.542 0 0 0-4.542-4.543"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.43 12.03a4.542 4.542 0 0 0 4.542-4.543"/>`),
		g.Group(children),
	)
}
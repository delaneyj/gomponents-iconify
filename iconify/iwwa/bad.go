package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20 38.5C9.799 38.5 1.5 30.201 1.5 20S9.799 1.5 20 1.5S38.5 9.799 38.5 20S30.201 38.5 20 38.5zm0-36C10.351 2.5 2.5 10.351 2.5 20c0 9.649 7.851 17.5 17.5 17.5S37.5 29.649 37.5 20c0-9.649-7.851-17.5-17.5-17.5z"/><path fill="currentColor" d="m11.477 29.011l-.851-.525a10.975 10.975 0 0 1 9.389-5.235c3.824 0 7.321 1.937 9.355 5.18l-.847.531a9.991 9.991 0 0 0-8.508-4.711a9.978 9.978 0 0 0-8.538 4.76z"/><circle cx="14.666" cy="15.309" r="1" fill="currentColor"/><circle cx="25.333" cy="15.309" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}
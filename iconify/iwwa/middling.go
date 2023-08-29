package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Middling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20 38.5C9.799 38.5 1.5 30.201 1.5 20S9.799 1.5 20 1.5S38.5 9.799 38.5 20S30.201 38.5 20 38.5zm0-36C10.351 2.5 2.5 10.351 2.5 20c0 9.649 7.851 17.5 17.5 17.5S37.5 29.649 37.5 20c0-9.649-7.851-17.5-17.5-17.5z"/><circle cx="14.667" cy="15.309" r="1" fill="currentColor"/><circle cx="25.334" cy="15.309" r="1" fill="currentColor"/><path fill="currentColor" d="M12.792 25.309h14.416v1H12.792z"/>`),
		g.Group(children),
	)
}
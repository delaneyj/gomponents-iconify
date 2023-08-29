package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Good(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20 38.5C9.799 38.5 1.5 30.201 1.5 20S9.799 1.5 20 1.5S38.5 9.799 38.5 20S30.201 38.5 20 38.5zm0-36C10.351 2.5 2.5 10.351 2.5 20c0 9.649 7.851 17.5 17.5 17.5c9.649 0 17.5-7.851 17.5-17.5c0-9.649-7.851-17.5-17.5-17.5z"/><path fill="currentColor" d="M20.016 31.399c-3.851 0-7.36-1.957-9.389-5.235l.851-.525a9.98 9.98 0 0 0 8.538 4.761a9.988 9.988 0 0 0 8.508-4.711l.848.531a10.987 10.987 0 0 1-9.356 5.179z"/><circle cx="14.416" cy="13.582" r="1" fill="currentColor"/><circle cx="25.084" cy="13.582" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}
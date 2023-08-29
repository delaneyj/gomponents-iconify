package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M21.049 14.741c3.58.43 8.219 4.52 8.6 8.28c0 0 9.449-9.471 9.439-9.5c1.939-1.94 1.87-2.75-.15-4.561l-4.08-3.68c-1.948-1.76-2.698-1.69-4.528.13l-9.281 9.331zm4.95 8.321c-.3-1.351-2.929-4.512-5.759-4.98c-7.1-1.83-13.89 5.88-13.49 11.491c.641 4.581-5.35 5.22-6.25 5.43c5.92 2.062 8.439 2.541 13.919 1.38c4.78-1.009 14.03-5.71 11.58-13.321z"/>`),
		g.Group(children),
	)
}
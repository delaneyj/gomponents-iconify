package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.625 2.655A9 9 0 0 1 19 11a1 1 0 1 1-2 0a7 7 0 0 0-9.625-6.492a1 1 0 1 1-.75-1.853ZM4.662 4.959A1 1 0 0 1 4.75 6.37A6.97 6.97 0 0 0 3 11a1 1 0 1 1-2 0a8.97 8.97 0 0 1 2.25-5.953a1 1 0 0 1 1.412-.088Z"/><path d="M5 11a5 5 0 1 1 10 0a1 1 0 1 1-2 0a3 3 0 1 0-6 0c0 1.677-.345 3.276-.968 4.729a1 1 0 1 1-1.838-.789A9.964 9.964 0 0 0 5 11Zm8.921 2.012a1 1 0 0 1 .831 1.145a19.86 19.86 0 0 1-.545 2.436a1 1 0 1 1-1.92-.558c.207-.713.371-1.445.49-2.192a1 1 0 0 1 1.144-.83Z"/><path d="M10 10a1 1 0 0 1 1 1c0 2.236-.46 4.368-1.29 6.304a1 1 0 0 1-1.838-.789A13.952 13.952 0 0 0 9 11a1 1 0 0 1 1-1Z"/></g>`),
		g.Group(children),
	)
}
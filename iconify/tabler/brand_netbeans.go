package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNetbeans(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.875 6.27A2.225 2.225 0 0 1 21 8.218v7.284c0 .809-.443 1.555-1.158 1.948l-6.75 4.27a2.269 2.269 0 0 1-2.184 0l-6.75-4.27A2.225 2.225 0 0 1 3 15.502V8.217c0-.809.443-1.554 1.158-1.947l6.75-3.98a2.33 2.33 0 0 1 2.25 0l6.75 3.98h-.033z"/><path d="M15.5 9.43a1 1 0 0 1 .5.874v3.268a1 1 0 0 1-.515.874l-3 1.917a1 1 0 0 1-.97 0l-3-1.917A1 1 0 0 1 8 13.573v-3.269a1 1 0 0 1 .514-.874l3-1.786c.311-.173.69-.173 1 0l3 1.787H15.5z"/><path d="M12 21v-9L4.5 7.5M12 12l7.5-4.5M12 3v4.5m7.5 8.5L16 14m-8 0l-3.5 2"/></g>`),
		g.Group(children),
	)
}
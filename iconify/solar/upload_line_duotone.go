package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M17 9.002c2.175.012 3.353.109 4.121.877C22 10.758 22 12.172 22 15v1c0 2.829 0 4.243-.879 5.122C20.243 22 18.828 22 16 22H8c-2.828 0-4.243 0-5.121-.878C2 20.242 2 18.829 2 16v-1c0-2.828 0-4.242.879-5.121c.768-.768 1.946-.865 4.121-.877" opacity=".5"/><path stroke-linejoin="round" d="M12 15V2m0 0l3 3.5M12 2L9 5.5"/></g>`),
		g.Group(children),
	)
}
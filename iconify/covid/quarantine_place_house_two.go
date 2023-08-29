package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuarantinePlaceHouseTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.017.768H8.983m1.517 0v2.6m-5.822-.817L3.606 3.623L2.533 4.696m1.073-1.073l1.838 1.839M.75 9.001v3.033m0-1.516h2.6m-.817 5.821l1.073 1.073l1.072 1.072m-1.072-1.072l1.838-1.839M18.467 4.696l-1.073-1.073l-1.072-1.072m1.072 1.072l-1.838 1.839M8.983 9.868a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.229 6.24a7.149 7.149 0 1 0-7.762 11.128m9.141-7.611a7.128 7.128 0 0 0-.253-1.272v0c-.24-.81-.62-1.57-1.126-2.245m-3.479 9.69v5.8a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5v-5.8"/><path stroke-linecap="round" stroke-linejoin="round" d="m11.25 17.243l4.518-3.954a2.252 2.252 0 0 1 2.964 0l4.518 3.954m-4.5 5.989h-3v-3a1.5 1.5 0 1 1 3 0v3Z"/><path d="M9.358 13.368a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		g.Group(children),
	)
}
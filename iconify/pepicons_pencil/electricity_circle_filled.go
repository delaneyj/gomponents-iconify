package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricityCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilElectricityCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M18 11.5h-3.813l2.273-5.303A.5.5 0 0 0 16 5.5h-5a.5.5 0 0 0-.46.303l-3 7A.5.5 0 0 0 8 13.5h2.474l-2.938 7.314c-.2.497.417.918.807.55l5.024-4.743l4.958-4.241A.5.5 0 0 0 18 11.5Zm-4.571 1h3.217l-3.948 3.378l-3.385 3.195l2.365-5.887a.5.5 0 0 0-.464-.686H8.758l2.572-6h3.912l-2.273 5.303a.5.5 0 0 0 .46.697Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilElectricityCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
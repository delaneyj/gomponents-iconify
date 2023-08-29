package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricityCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M15.187 13H19a.5.5 0 0 1 .325.88l-4.958 4.241l-5.024 4.743c-.39.368-1.007-.053-.807-.55L11.474 15H9a.5.5 0 0 1-.46-.697l3-7A.5.5 0 0 1 12 7h5a.5.5 0 0 1 .46.697L15.187 13Z" opacity=".2"/><path fill-rule="evenodd" d="M18 11.5h-3.813l2.273-5.303A.5.5 0 0 0 16 5.5h-5a.5.5 0 0 0-.46.303l-3 7A.5.5 0 0 0 8 13.5h2.474l-2.938 7.314c-.2.497.417.918.807.55l5.024-4.743l4.958-4.241A.5.5 0 0 0 18 11.5Zm-4.571 1h3.217l-3.948 3.378l-3.385 3.195l2.365-5.887a.5.5 0 0 0-.464-.686H8.758l2.572-6h3.912l-2.273 5.303a.5.5 0 0 0 .46.697Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
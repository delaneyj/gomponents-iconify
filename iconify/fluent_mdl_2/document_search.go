package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 896q79 0 149 30t122 82t82 123t30 149q0 80-30 149t-82 122t-122 83t-149 30q-60 0-117-18t-105-53l-437 436q-19 19-45 19t-45-19t-19-45q0-26 19-45l436-437q-35-48-53-105t-18-117q0-79 30-149t82-122t122-83t150-30zm0 640q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20zm-656 128l-128 128H256V0h859l549 549v219h-128V640h-512V128H384v1536h624zm144-1152h293l-293-293v293z"/>`),
		g.Group(children),
	)
}
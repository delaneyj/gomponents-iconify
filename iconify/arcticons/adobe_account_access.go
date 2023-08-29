package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeAccountAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 35.5c5.333-1.839 9.134-6.112 9.73-11.552a46.804 46.804 0 0 0-.041-7.093a1.499 1.499 0 0 0-1.455-1.396c-2.413-.074-5.217-.476-7.384-2.61a1.21 1.21 0 0 0-1.7 0c-2.166 2.134-4.97 2.536-7.383 2.61a1.499 1.499 0 0 0-1.455 1.396a46.79 46.79 0 0 0-.043 7.093c.597 5.44 4.398 9.714 9.731 11.552Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.44 22.792h7.12a.49.49 0 0 1 .489.489v5.262a.49.49 0 0 1-.49.489H20.44a.489.489 0 0 1-.489-.49v-5.261c0-.27.219-.49.489-.49h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.32 22.792v-1.123a2.678 2.678 0 0 1 5.355 0v1.123"/>`),
		g.Group(children),
	)
}
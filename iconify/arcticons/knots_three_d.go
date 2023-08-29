package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnotsThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.911 21.68a22.84 22.84 0 0 0 7.446-4.99c4.17-4.17 6.868-10.792 6.868-10.792m-6.847 24.914c-.701 2.812-1.212 4.75-2.424 6.09c-2.797 3.09-8.878 4.823-12.755-1.03c-2.502-3.778-1.177-8.781 2.404-11.48a8.333 8.333 0 0 1 1.5-.99M43.5 12.766a51.683 51.683 0 0 0-7.678 7.015c-1.84 2.01-2.607 4.065-3.392 6.98m-19.22 7.324A88.4 88.4 0 0 0 4.5 39.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.3 18.669c1.493 3.147 1.098 6.01-.78 7.49c-4.13 3.251-9.165 2.785-15.796 4.366a40.415 40.415 0 0 0-5.868 1.919M35.7 13.794a7.479 7.479 0 0 1 1.852 1.813"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.645 28.318c-2.248-4.7-2.721-9.593 1.58-13.933a7.31 7.31 0 0 1 7.43-2.113m2.038 29.83c-1.417-1.321-5.94-4.991-9.212-10.153"/>`),
		g.Group(children),
	)
}
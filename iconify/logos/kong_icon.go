package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KongIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#169FCC" d="m80.721 192.398l6.115-7.893h45.253l23.587 30.136l-4.193 10.225H93.126l1.398-10.225z"/><path fill="#14B59A" d="m94.559 89.641l21.851-37.749h25.462l114.129 132.566l-8.849 40.408h-48.938l3.07-11.343z"/><path fill="#1BC263" d="m121.082 44.72l10.462-19.348L162.932 0l53.782 42.164l-6.975 7.119l9.361 12.96v13.872l-26.799 21.903L147.33 44.72z"/><path fill="#16BDCC" d="M37.969 129.026h14.216l37.069-31.008l49.125 57.082l-13.856 20.789H79.177l-31.311 39.816l-7.198 9.161H0v-48.801z"/>`),
		g.Group(children),
	)
}
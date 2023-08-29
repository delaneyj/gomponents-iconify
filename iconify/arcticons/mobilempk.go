package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobilempk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.54 39.076h4.742V43.5H8.54zm26.178 0h4.742V43.5h-4.742zM7.29 10.992h1.202v7.934a2.433 2.433 0 0 1-2.405-2.462v-4.241A1.218 1.218 0 0 1 7.29 10.99Zm33.42 0h-1.202v7.934a2.433 2.433 0 0 0 2.405-2.462v-4.241a1.218 1.218 0 0 0-1.202-1.232ZM24 4.5c-5.511 0-15.46.856-15.46 3.097v31.479h30.921V7.596C39.461 5.357 29.439 4.5 24 4.5Zm.024 4.232q8.185 0 12.197 1.224a.961.961 0 0 1 .667.926v17.019a.975.975 0 0 1-.975.975H12.135a.975.975 0 0 1-.975-.975V10.882a.961.961 0 0 1 .667-.926q4.014-1.224 12.197-1.224Zm-10.435 22.77a2.502 2.502 0 0 1 0 5h0a2.465 2.465 0 0 1-2.429-2.5h0a2.465 2.465 0 0 1 2.429-2.5Zm20.87 0a2.502 2.502 0 0 1 0 5h0a2.465 2.465 0 0 1-2.428-2.5h0a2.465 2.465 0 0 1 2.429-2.5Z"/>`),
		g.Group(children),
	)
}
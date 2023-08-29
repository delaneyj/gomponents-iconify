package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningWaterAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 40a2 2 0 0 0 2-2c0-1.066-.654-2.37-1.59-3.6c-.138-.18-.276-.351-.41-.512c-.134.16-.272.332-.41.512C22.655 35.63 22 36.934 22 38a2 2 0 0 0 2 2Zm0 2a4 4 0 0 0 4-4c0-3.5-4-7-4-7s-4 3.5-4 7a4 4 0 0 0 4 4Zm-2-30a2 2 0 1 1 4 0v12h-4V12Zm-1 14v2h6v-2h-6Zm8 1v3H19v-3H4v-2h1v-3a2 2 0 0 1 2-2h3v-1H7v-5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v5h-3v1h5v-8a4 4 0 0 1 8 0v8h5v-1h-3v-5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v5h-3v1h3a2 2 0 0 1 2 2v3h1v2H29Zm12-2v-3H28v2h1v1h12Zm-5-5v-1h-1v1h1Zm-17 4v1H7v-3h13v2h-1Zm-6-4h-1v-1h1v1Zm3-6H9v3h1v-2h2v2h1v-2h2v2h1v-3Zm16 0h7v3h-1v-2h-2v2h-1v-2h-2v2h-1v-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
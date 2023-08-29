package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgioIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><radialGradient id="logosEdgioIcon0" cx="104.362%" cy="13.088%" r="94.575%" fx="104.362%" fy="13.088%"><stop offset="0%" stop-color="#01B07D"/><stop offset="100%" stop-color="#01B07D" stop-opacity="0"/></radialGradient><radialGradient id="logosEdgioIcon1" cx="68.749%" cy="120.916%" r="68.487%" fx="68.749%" fy="120.916%"><stop offset="0%" stop-color="#00AAE5"/><stop offset="100%" stop-color="#00AAE5" stop-opacity="0"/></radialGradient><linearGradient id="logosEdgioIcon2" x1="3.185%" x2="57.325%" y1="2.866%" y2="58.917%"><stop offset="0%" stop-color="#793092"/><stop offset="100%" stop-color="#6144A1"/></linearGradient></defs><path fill="url(#logosEdgioIcon2)" d="M0 0h256v256H0z"/><path fill="url(#logosEdgioIcon0)" d="M0 0h256v256H0z"/><path fill="url(#logosEdgioIcon1)" d="M0 0h256v256H0z"/><path fill="#FFF" d="m170.908 77.201l16.225-29.253H68.867v160.103h118.266l-16.144-29.293h-70.24v-36.819h55.755l8.983-29.294h-64.738V77.201z"/>`),
		g.Group(children),
	)
}
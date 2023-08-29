package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rugby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 0q106 0 199 40t163 109t110 163t40 200q0 212-55 408t-155 367t-240 311t-311 240t-367 155t-408 55q-106 0-199-40t-163-109t-110-163t-40-200q0-212 55-408t155-367t240-311t311-240t367-155t408-55zm0 128q-194 0-373 50T827 320T541 541T321 826t-142 336t-51 374q0 60 18 117t53 105L1758 199q-48-35-105-53t-117-18zM512 1920q194 0 373-50t336-142t286-221t220-285t142-336t51-374q0-60-18-117t-53-105L290 1849q48 35 105 53t117 18z"/>`),
		g.Group(children),
	)
}
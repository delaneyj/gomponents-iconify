package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M482 277q8 0 14.5 7t6.5 15v85q0 8-6.5 14.5T482 405H375q-8 0-14.5-6.5T354 384v-85q0-8 6.5-15t14.5-7v-32q0-22 15.5-37.5t38-15.5t38 15.5T482 245v32zm-21 0v-32q0-12-9.5-22T429 213t-22.5 10t-9.5 22v32h64zm-139-32v56l-75 94L0 85Q114 0 247.5 0T495 85l-45 56q-6-2-21-2q-45 0-76 31t-31 75z"/>`),
		g.Group(children),
	)
}
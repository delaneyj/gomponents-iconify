package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosSettingsStrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M32 376h283.35c6.186-14.112 20.281-24 36.65-24s30.465 9.888 36.65 24H480v32h-91.35c-6.186 14.112-20.281 24-36.65 24s-30.465-9.888-36.65-24H32" fill="currentColor"/><path d="M32 240h91.35c6.186-14.112 20.281-24 36.65-24s30.465 9.888 36.65 24H480v32H196.65c-6.186 14.112-20.281 24-36.65 24s-30.465-9.888-36.65-24H32" fill="currentColor"/><path d="M32 104h283.35c6.186-14.112 20.281-24 36.65-24s30.465 9.888 36.65 24H480v32h-91.35c-6.186 14.112-20.281 24-36.65 24s-30.465-9.888-36.65-24H32" fill="currentColor"/>`),
		g.Group(children),
	)
}
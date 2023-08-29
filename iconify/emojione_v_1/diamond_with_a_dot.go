package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondWithADot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#25333a" d="M.002 31.803L31.984-.18l31.982 31.982l-31.982 31.982z"/><path fill="#724198" d="M60.51 28.355a4.874 4.874 0 0 1 0 6.895L35.42 60.34a4.876 4.876 0 0 1-6.894 0L3.436 35.25a4.875 4.875 0 0 1 0-6.892l25.09-25.09a4.873 4.873 0 0 1 6.893 0l25.09 25.09"/><path fill="#0867a3" d="m24.351 56.16l4.182 4.182a4.876 4.876 0 0 0 6.894 0l25.09-25.09a4.874 4.874 0 0 0 0-6.895l-25.09-25.09c14.376 25.297-4 46.19-11.08 52.887"/><path fill="#439dd7" d="M54.741 32.518L31.98 54.923L9.219 32.518L31.981 10.11z"/><path fill="#fff" d="m38.59 32.518l-6.609 6.502l-6.608-6.502l6.608-6.508z"/><path fill="#8dcbef" d="M31.981 10.11L21.706 32.518L31.98 54.923L9.219 32.518z"/>`),
		g.Group(children),
	)
}
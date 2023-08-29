package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOrPauseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M14.158 54.717c.42.17.895.283 1.368.283c.737 0 1.474-.283 2.106-.737l26.052-15.254l.263-.284C44.632 37.988 45 37.081 45 36.06s-.368-1.985-1.053-2.665l-.263-.284l-26.052-15.367c-.948-.794-2.316-.964-3.474-.454c-1.316.567-2.158 1.985-2.158 3.516V51.2c0 1.53.842 2.948 2.158 3.516zM60 17v38m-8-38v38"/>`),
		g.Group(children),
	)
}
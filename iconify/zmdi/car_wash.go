package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarWash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M298.5 91Q285 91 276 81.5T267 59q0-10 8-24.5T291 11l8-10q32 36 32 58q0 13-9.5 22.5t-23 9.5zM192 91q-13 0-22.5-9.5T160 59q0-10 8-24.5T184 11l8-10q32 36 32 58q0 13-9.5 22.5T192 91zM85.5 91q-13.5 0-23-9.5T53 59q0-10 8-24.5T77 11l8-10q32 36 32 58q0 13-9 22.5T85.5 91zM340 155l44 128v170q0 9-6.5 15.5T363 475h-22q-8 0-14.5-6.5T320 453v-21H64v21q0 9-6.5 15.5T43 475H21q-8 0-14.5-6.5T0 453V283l44-128q8-22 31-22h234q23 0 31 22zM74.5 368q13.5 0 23-9.5T107 336t-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5t9 22.5t22.5 9.5zm235 0q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T277 336t9.5 22.5t23 9.5zM43 261h298l-32-96H75z"/>`),
		g.Group(children),
	)
}
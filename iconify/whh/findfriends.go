package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Findfriends(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1005 1005q-19 19-45.5 19t-45.5-19L716 807q-120 89-268 89q-91 0-174-35.5T131 765T35.5 622T0 448t35.5-174T131 131t143-95.5T448 0t174 35.5T765 131t95.5 143T896 448q0 148-89 268l198 198q19 19 19 45.5t-19 45.5zM448 128q-87 0-160.5 43T171 287.5T128 448q0 121 81 212q56-33 143-46v-31q-96-58-96-240q0-74 54-112.5T448.5 192t138 38.5T640 343q0 190-96 243v28q87 12 144 46q80-91 80-212q0-87-43-160.5T608.5 171T448 128z"/>`),
		g.Group(children),
	)
}
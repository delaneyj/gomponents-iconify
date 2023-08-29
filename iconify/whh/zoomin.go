package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoomin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1005 1005q-19 19-45.5 19t-45.5-19L716 807q-120 89-268 89q-91 0-174-35.5T131 765T35.5 622T0 448t35.5-174T131 131t143-95.5T448 0t174 35.5T765 131t95.5 143T896 448q0 148-89 268l198 198q19 19 19 45.5t-19 45.5zM448 128q-87 0-160.5 43T171 287.5T128 448t43 161t116.5 116.5T448 768t160.5-42.5T725 609t43-161t-43-160.5T608.5 171T448 128zm128 384h-64v64q0 27-19 45.5T447.5 640t-45-18.5T384 576v-64h-64q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h64v-64q0-26 18.5-45t45-19t45.5 19t19 45v64h64q26 0 45 19t19 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}
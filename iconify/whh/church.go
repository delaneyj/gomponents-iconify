package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Church(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1017 576H449v64h480q13 0 22.5 9.5T961 672v320q0 13-9.5 22.5T929 1024H257V864q0-13-9.5-22.5T225 832t-22.5 9.5T193 864v160H33q-13 0-22.5-9.5T1 992V482q-1-16 8-27l184-177v-86H97q-13 0-22.5-9.5T65 160t9.5-22.5T97 128h96V32q0-13 9.5-22.5T225 0t22.5 9.5T257 32v96h96q13 0 22.5 9.5T385 160t-9.5 22.5T353 192h-96v86l177 170h476q16 0 27 16l84 96q5 7 3.5 11.5t-7.5 4.5zM833 864q0 13 9.5 22.5T865 896t22.5-9.5T897 864v-64q0-13-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5v64zm-128 0q0 13 9.5 22.5T737 896t22.5-9.5T769 864v-64q0-13-9.5-22.5T737 768t-22.5 9.5T705 800v64zm-128 0q0 13 9.5 22.5T609 896t22.5-9.5T641 864v-64q0-13-9.5-22.5T609 768t-22.5 9.5T577 800v64z"/>`),
		g.Group(children),
	)
}
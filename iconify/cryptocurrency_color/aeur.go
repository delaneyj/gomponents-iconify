package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aeur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><linearGradient id="cryptocurrencyColorAeur0" x1="50%" x2="50%" y1="0%" y2="143.239%"><stop offset="0%" stop-color="#FFF"/><stop offset="3%" stop-color="#FFF" stop-opacity=".83"/><stop offset="7%" stop-color="#FFF" stop-opacity=".66"/><stop offset="11%" stop-color="#FFF" stop-opacity=".5"/><stop offset="15%" stop-color="#FFF" stop-opacity=".37"/><stop offset="19%" stop-color="#FFF" stop-opacity=".25"/><stop offset="25%" stop-color="#FFF" stop-opacity=".16"/><stop offset="30%" stop-color="#FFF" stop-opacity=".09"/><stop offset="37%" stop-color="#FFF" stop-opacity=".04"/><stop offset="47%" stop-color="#FFF" stop-opacity=".01"/><stop offset="100%" stop-color="#FFF" stop-opacity="0"/></linearGradient></defs><g fill="none"><circle cx="16" cy="16" r="16" fill="#051D2D"/><g fill="url(#cryptocurrencyColorAeur0)" transform="translate(9 6)"><path d="M6.993 13.986a6.993 6.993 0 1 1 6.993-6.993a7.002 7.002 0 0 1-6.993 6.993zM7 6.951A.049.049 0 1 0 7.049 7a.055.055 0 0 0-.05-.05z"/><path d="M6.993 20.986a6.993 6.993 0 1 1 6.993-6.993a7.002 7.002 0 0 1-6.993 6.993zM7 13.951a.049.049 0 1 0 .049.049a.055.055 0 0 0-.05-.05z"/></g></g>`),
		g.Group(children),
	)
}
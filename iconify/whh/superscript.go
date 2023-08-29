package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 192H832v64h160q13 0 22.5 9.5t9.5 22.5t-9.5 22.5T992 320H800q-12 0-20.5-7.5T769 294V154q2-11 10.5-18.5T800 128h160V64H800q-13 0-22.5-9.5T768 32t9.5-22.5T800 0h192q12 0 20.5 7.5T1023 26v140q-2 11-10.5 18.5T992 192zm-369.5 814.5Q606 1023 583 1023t-40-17L320 729L96 1006q-16 17-39.5 17t-40-16.5t-16.5-40T17 927l231-288L17 352Q0 335 0 312t16.5-39.5t40-16.5T96 272l224 278l223-278q17-16 40-16t39.5 16.5T639 312t-16 40L392 639l231 288q16 16 16 39.5t-16.5 40z"/>`),
		g.Group(children),
	)
}
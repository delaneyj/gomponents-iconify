package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiSake0" d="M34 66.636c10 .364 29.7 0 31 0C64 52 64 37 58 24M37.547 38.248c1.15-5.812 2.64-10.806 4.379-14.662l.207-.374m-4.586 15.036a121.14 121.14 0 0 0-1.266 7.788M18.04 51h16.33c0 6.952-2.286 12.776-7.116 15.68h-14.4C8.023 63.776 5.736 57.952 5.736 51h12.305zM64.37 8c0 6.778-2.173 12.483-6.758 15.456m-15.635-.354c-4.23-3.059-6.24-8.584-6.24-15.102"/></defs><g fill="#FFF"><path d="M65 66.64c-1.3 0-22 .36-32 0C34 52 36 37 41 24l1-.82c-4.32-3.03-6.37-8.6-6.37-16.18h28.74c0 7.58-2.05 13.15-6.37 16.18V24c6 13 6 28 7 42.64z"/><use href="#openmojiSake0"/></g><use href="#openmojiSake0" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}
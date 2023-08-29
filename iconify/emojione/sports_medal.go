package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMedal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#428bc1" d="M47.8 2L35.2 21h6.3L54.1 2z"/><path fill="#e8e8e8" d="M41.5 2L28.8 21h6.4L47.8 2z"/><path fill="#ed4c5c" d="M35.2 2L22.5 21h6.3L41.5 2z"/><path fill="#ffc200" d="M20.4 16.8c-.6 0-1.1.5-1.1 1.1v9.5c0 .6.5 1.1 1.1 1.1h23.2c.6 0 1.1-.5 1.1-1.1v-9.5c0-.6-.5-1.1-1.1-1.1H20.4m22.1 7.3c0 .6-.5 1.1-1.1 1.1h-19c-.6 0-1.1-.5-1.1-1.1v-4.2c0-.6.5-1.1 1.1-1.1h19c.6 0 1.1.5 1.1 1.1v4.2"/><path fill="#ed4c5c" d="M22.5 21h6.3L16.2 2H9.9z"/><path fill="#e8e8e8" d="M28.8 21h6.4L22.5 2h-6.3z"/><path fill="#3e4347" d="m33.1 5.2l-3.2 4.7L37.3 21h4.2l1-1.6z" opacity=".5"/><path fill="#428bc1" d="M35.2 21h6.3L28.8 2h-6.3z"/><circle cx="32" cy="42.3" r="19.7" fill="#ffc200"/><path fill="#e68a00" d="M32.3 24.4c-10.1 0-18.2 8.2-18.2 18.2c0 3 .7 5.8 2 8.3c-.6-2-1-4.1-1-6.3c0-10.7 8.2-19.4 18.7-20.2h-1.5"/><path fill="#ffe394" d="M46 31c5.1 9 2.5 20.6-6.4 26.5c-1.8 1.2-3.8 2.1-5.8 2.7c2.8-.3 5.5-1.3 8-3c8.4-5.6 10.6-16.8 5.1-25L46 31"/><path fill="#f2b200" d="M32 34.3v-6.4l-3.2 10l1.4 1.8z"/><path fill="#e68a00" d="m33.8 39.7l1.4-1.8l-3.2-10v6.4z"/><path fill="#c47500" d="m34.8 43l2.4 1.1l8.5-6.2l-6.3 1.8z"/><path fill="#ffe394" d="m39.4 39.7l6.3-1.8H35.2l-1.4 1.8z"/><path fill="#ffd252" d="m30.2 39.7l-1.4-1.8H18.3l6.3 1.8z"/><path fill="#ffdb75" d="m24.6 39.7l-6.3-1.8l8.4 6.2l2.5-1.1z"/><path fill="#e68a00" d="m34.8 43l1.8 5.4l3.9 5.7l-3.3-10z"/><path fill="#f2b200" d="M32 45.1v2.8l8.5 6.2l-3.9-5.7zM29.2 43l-2.5 1.1l-3.2 10l3.9-5.7z"/><path fill="#e68a00" d="m27.4 48.4l-3.9 5.7l8.5-6.2v-2.8z"/><path fill="#ffce31" d="M33.8 39.7L32 34.3l-1.8 5.4h-5.6l4.6 3.3l-1.8 5.4l4.6-3.3l4.6 3.3l-1.8-5.4l4.6-3.3z"/>`),
		g.Group(children),
	)
}
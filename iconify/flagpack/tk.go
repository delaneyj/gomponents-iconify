package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTk0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTk2)"><use href="#flagpackTk0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackTk1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackTk1)"><path fill="#F7FCFF" d="m3.294 10.063l-.885.548l.344-.92L2 8.995h.913L3.294 8l.292.995H4.5l-.659.696l.323.92l-.87-.548Zm4-4l-.885.548l.344-.92L6 4.995h.913L7.294 4l.292.995H8.5l-.659.696l.323.92l-.87-.548Zm4 4l-.885.548l.344-.92L10 8.995h.913L11.294 8l.292.995h.914l-.659.696l.323.92l-.87-.548Zm-4 4l-.885.548l.344-.92L6 12.995h.913L7.294 12l.292.995H8.5l-.659.697l.323.92l-.87-.549Z"/><path fill="#FECA00" d="M24.842 5.463c-4.083 2.017-17.5 11.082-17.5 11.082h22.35c-.223-.05-.431-.091-.628-.13c-1.64-.327-2.45-.488-4.222-4.018c-1.984-3.952 0-6.934 0-6.934ZM6.616 17.965l-.306.76l.306.835l22.753.44l.631-1.16l-.63-.835l-22.754-.04Z"/></g></g><defs><clipPath id="flagpackTk2"><use href="#flagpackTk0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroSlrCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMicroSlrCamera0"><g fill="none"><path fill="#555" d="M44 26c0 2.64-.73 5.108-2 7.215A14.073 14.073 0 0 1 37.215 38c-2.107 1.27-4.576 2-7.215 2c-2.64 0-5.108-.73-7.215-2C18.719 35.55 16 31.093 16 26s2.72-9.55 6.785-12c2.107-1.27 4.576-2 7.215-2c2.64 0 5.108.73 7.215 2A14.073 14.073 0 0 1 42 18.785c1.27 2.107 2 4.576 2 7.215Z"/><path fill="#555" d="M4 14v24h18.785C18.719 35.55 16 31.093 16 26s2.72-9.55 6.785-12H4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37.215 14c-2.107-1.27-4.576-2-7.215-2c-2.64 0-5.108.73-7.215 2m14.43 0H42v4.785M37.215 14A14.073 14.073 0 0 1 42 18.785M22.785 14H4v24h18.785m0-24C18.719 16.45 16 20.907 16 26s2.72 9.55 6.785 12m14.43 0c-2.107 1.27-4.576 2-7.215 2c-2.64 0-5.108-.73-7.215-2m14.43 0H42v-4.785M37.215 38A14.073 14.073 0 0 0 42 33.215m0-14.43c1.27 2.107 2 4.576 2 7.215c0 2.64-.73 5.108-2 7.215"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 9h9v5H8z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 22v8m26-4a6 6 0 0 1-6 6m-6-6a6 6 0 0 1 6-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMicroSlrCamera0)"/>`),
		g.Group(children),
	)
}
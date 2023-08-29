package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FriendsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFriendsCircle0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="#555" fill-rule="evenodd" d="M31 7v17V7Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 7v17"/><path fill="#555" fill-rule="evenodd" d="m16.636 6.636l14.142 14.142L16.636 6.636Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m16.636 6.636l14.142 14.142"/><path fill="#555" fill-rule="evenodd" d="M7 17h17H7Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 17h17"/><path fill="#555" fill-rule="evenodd" d="M20.364 17.636L6.222 31.778l14.142-14.142Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20.364 17.636L6.222 31.778"/><path fill="#555" fill-rule="evenodd" d="M17 25v17v-17Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 25v17"/><path fill="#555" fill-rule="evenodd" d="m17.636 27.636l14.142 14.142l-14.142-14.142Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17.636 27.636l14.142 14.142"/><path fill="#555" fill-rule="evenodd" d="M24 31h18h-18Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 31h18"/><path fill="#555" fill-rule="evenodd" d="M42.364 16.636L28.222 30.778l14.142-14.142Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42.364 16.636L28.222 30.778"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 31a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFriendsCircle0)"/>`),
		g.Group(children),
	)
}
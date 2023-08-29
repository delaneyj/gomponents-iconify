package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FriendsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path fill="#2F88FF" fill-rule="evenodd" d="M31 7V24V7Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 7V24"/><path fill="#2F88FF" fill-rule="evenodd" d="M16.6357 6.63599L30.7779 20.7781L16.6357 6.63599Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16.6357 6.63599L30.7779 20.7781"/><path fill="#2F88FF" fill-rule="evenodd" d="M7 17H24H7Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 17H24"/><path fill="#2F88FF" fill-rule="evenodd" d="M20.3643 17.636L6.22212 31.7781L20.3643 17.636Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20.3643 17.636L6.22212 31.7781"/><path fill="#2F88FF" fill-rule="evenodd" d="M17 25V42V25Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 25V42"/><path fill="#2F88FF" fill-rule="evenodd" d="M17.6357 27.636L31.7779 41.7781L17.6357 27.636Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17.6357 27.636L31.7779 41.7781"/><path fill="#2F88FF" fill-rule="evenodd" d="M24 31L42 31L24 31Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 31L42 31"/><path fill="#2F88FF" fill-rule="evenodd" d="M42.3643 16.636L28.2221 30.7781L42.3643 16.636Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42.3643 16.636L28.2221 30.7781"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 31C27.866 31 31 27.866 31 24C31 20.134 27.866 17 24 17C20.134 17 17 20.134 17 24C17 27.866 20.134 31 24 31Z"/></g>`),
		g.Group(children),
	)
}
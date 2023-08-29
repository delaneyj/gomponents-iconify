package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailBill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 0H64Q37 0 18.5 18.5T0 64v256q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320V64q0-27-18.5-45.5T448 0zm21 320q0 21-21 21H64q-21 0-21-21V64q0-21 21-21h21v85l43-21l43 21V43h277q21 0 21 21v256zM405 64h-42q-10 0-16 6t-6 15v64q0 22 22 22h42q22 0 22-22V85q0-9-6-15t-16-6zm-21 149H277q-17 0-29.5 13T235 256v21q0 18 12.5 30.5T277 320h107q17 0 30-12.5t13-30.5v-21q0-17-13-30t-30-13zm0 64H277v-21h107v21z"/>`),
		g.Group(children),
	)
}
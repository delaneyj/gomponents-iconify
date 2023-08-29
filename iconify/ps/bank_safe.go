package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankSafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 0H64Q37 0 18.5 18.5T0 64v405q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V64q0-27-18.5-45.5T448 0zm21 469H43V64q0-21 21-21h384q21 0 21 21v405zM384 64H128q-17 0-30 12.5T85 107v42H64v43h21v128H64v43h21v42q0 18 13 30.5t30 12.5h256q27 0 45.5-18.5T448 384V128q0-27-18.5-45.5T384 64zm21 320q0 21-21 21H128v-42h21v-43h-21V192h21v-43h-21v-42h256q21 0 21 21v256zm-70-100q6-13 6-28q0-20-12-36.5T299 196v-47q0-9-6-15t-16-6q-9 0-15 6t-6 15v47q-43 15-43 60q0 15 7 28l-43 42q-14 16 0 30q6 7 15 7t15-7l43-42q12 6 27 6q16 0 28-6l43 42q6 7 15 7q8 0 15-7q14-14 0-30zm-58-7q-21 0-21-21t21-21q22 0 22 21t-22 21z"/>`),
		g.Group(children),
	)
}
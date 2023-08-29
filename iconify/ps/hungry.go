package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hungry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 171q0-47-23-86t-62-62v-2h-2Q319 0 277 0h-42q-39 0-84 21h-2v2q-39 23-62 62t-23 86q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171zM64 282q-21-18-21-47q0-22 21-22v69zm341 38q0 62-43.5 105.5T256 469t-105.5-43.5T107 320V171q0-57 42-96v10q94 38 214 0V75q42 39 42 96v149zm43-38v-69q21 0 21 22q0 29-21 47zm-171 81h-42q-22 0-22 21t22 21h42q22 0 22-21t-22-21zm-42-86h-43q0 18 12.5 30.5T235 320v-43zm42 43q18 0 30.5-12.5T320 277h-43v43zm34-90q6-4 9-4v9q0 9 6 15t15 6q10 0 16-6t6-15v-22h21v-42q-69 0-100 25zm-119-4q2 1 4.5 2t4.5 2l25-34q-35-25-98-25v42q13 0 21 2v20q0 9 6 15t16 6q9 0 15-6t6-15v-9z"/>`),
		g.Group(children),
	)
}
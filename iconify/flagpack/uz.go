package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackUz0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackUz1)"><use href="#flagpackUz0"/><path fill="#14DC5A" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#0099B5" fill-rule="evenodd" d="M0 0h32v10H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" stroke="#C51918" d="M-2 9.5h-.5v7h37v-7H-2Z"/><path fill="#F7FCFF" fill-rule="evenodd" d="m14.541 3.006l-.673.374l.192-.76l-.644-.558h.842l.282-.72l.331.72h.718l-.564.558l.271.76l-.755-.374Zm-3.608 2.291l.673-.374l.755.374l-.27-.76l.563-.558h-.718l-.33-.72l-.283.72h-.842l.644.558l-.192.76ZM8.428 6.961l-.673.374l.192-.76l-.645-.559h.842l.283-.719l.33.72h.719l-.564.558l.27.76l-.754-.374Zm3.178 0l-.673.374l.192-.76l-.644-.559h.842l.282-.719l.331.72h.718l-.564.558l.271.76l-.755-.374Zm2.98 0l-.673.374l.192-.76l-.644-.559h.842l.282-.719l.331.72h.718l-.564.558l.271.76l-.755-.374Zm-.673-1.664l.673-.374l.755.374l-.27-.76l.563-.558h-.718l-.33-.72l-.283.72h-.842l.644.558l-.192.76Zm3 2.038l.672-.374l.756.374l-.271-.76l.564-.559h-.718l-.332-.719l-.282.72h-.842l.645.558l-.193.76Zm.672-2.412l-.673.374l.193-.76l-.645-.558h.842l.282-.72l.332.72h.718l-.564.558l.27.76l-.755-.374Zm-.717-1.543l.673-.374l.755.374l-.271-.76l.564-.558h-.718l-.331-.72l-.283.72h-.842l.645.558l-.192.76Zm3.748 3.581l-.673.374l.192-.76l-.645-.559h.842l.283-.719l.33.72h.719l-.564.558l.27.76l-.754-.374Zm-.673-1.664l.673-.374l.755.374l-.271-.76l.564-.558h-.718l-.331-.72l-.283.72h-.842l.645.558l-.192.76Zm.628-2.29l-.673.373l.192-.76l-.645-.558h.842l.283-.72l.331.72h.718l-.564.558l.271.76l-.755-.374ZM5.885 8.24s-2.416-.656-2.37-3.08C3.56 2.739 6 2.11 6 2.11c-.997-.377-3.945.13-4 3.028c-.054 2.9 2.956 3.47 3.885 3.104Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackUz1"><use href="#flagpackUz0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
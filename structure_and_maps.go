package main

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) updatecontact2(newemail string) (p2 person) { // This is a value receiver it won't update the email
	p.contact.email = newemail
	return p
}

func (p person) updatecontact(newemail string) { // This is a value receiver it will update the email
	p.contact.email = newemail

}

func (p *person) updatecontactpointer(newemail string) {
	(*p).contact.email = newemail
}

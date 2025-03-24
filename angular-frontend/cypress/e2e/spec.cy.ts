describe('Test1: Login form', () => {
  it('Visits the login project page and inputs', () => {
    cy.visit('/login')
    cy.get('input[name="email"]').type('test@mail.com', {force: true}).blur(); 
    cy.get('input[name="password"]').type('test', {force: true}).blur(); 
    
    cy.get('button').contains('Login').click();
    
    cy.on('window:alert', (str) => {
     expect(str).to.equal('Login successful!')   
    })
    
  })
})

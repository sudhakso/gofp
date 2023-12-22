# gofp - Design Patterns in Golang

## Principles

Few principles to take into consideration while doing code designing (SOLID)

1. SRP -  Single Responsibility Principle

   :heavy_check_mark: An Object can have only one primary responsibility and hence would change when the primary responsibility changes.

   An easy example woud be,
   * Journal type, with Add, Remove and String functions
   * Persistence type, with SaveToFile() implementation

   When compared to 
   * Journal type, with Add, Remove, String and SaveToFile() implementation

   It is also called as *SoC* : Separation of concerns.
   
2. OCP - Open Closed Principle

3. LSP - Liskov Substitution Principle

4. ISP - Interface Segregation Principle

5. DIP - Dependency Inversion Principle
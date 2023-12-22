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

    :heavy_check_mark: A type is closed for *any* modifications to suit the needs of future requirements but is open for *extensions*.

    An easy similarity would be to compare it to database design. 
    * You do not change the existing fields in the DB table
    * You add new tables, and have new features exposed through application code.

    Another concrete example is,
    * *Product* type, with *Size* & *Color* as its attributes
    * Extending the Product type to support *filtering* capabilities, such that, you can filter the Products by *Size*, *Colo* or both.

    When compared to
    * Implement *Filter* type, to support filtering by *Size*, *Color*, both *size* & *color*, and in future keep changing the *Filter* type to add new features.

3. LSP - Liskov Substitution Principle

    :heavy_check_mark: If a type implements an interface, all its inherited sub-types can be substituted by the client in their implementation and everything should work as expected.

    An easy example would be,
    * A rectangle has four sides, and implements *Dimension* interface with attributes,
        * length (with GET/SET functions)
        * breadth (with GET/SET functions)
    * A square *is a* rectange with length == breadth, and implements *Dimension* interface and has attributes,
        * size
    So, if *Dimension* interface has following functions,
    * getWidth(), setWidth()
    * getHeight(), setHeight()
    Eventhough *Square* implements Dimension interface, it cannot expose setWidth & setHeight functions separately. So eventhoug *Square* is a *Rectangle*, it cannot be implemented through inheritence. 

4. ISP - Interface Segregation Principle

5. DIP - Dependency Inversion Principle
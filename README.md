# Balabola GO Client
GO client for Yandex balabola service 

### Example 
```go
	ctx := context.TODO()

	c := balabola.New(ctx)

	nonsense, err := c.TalkNonsense("Акции bitop выросли на 5%", TextStyleWithoutStyle, 1)
	if err != nil {
		log.Default().Print(err)
		
		return err
	}
	
	fmt.Println(nonsense)
```
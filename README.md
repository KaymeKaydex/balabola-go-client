# Balabola GO Client
GO client for Yandex balabola service 

### Yandex warning
The neural network does not know what it is saying, and can say anything — if anything, do not be offended. When distributing the resulting texts, remember about responsibility.
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

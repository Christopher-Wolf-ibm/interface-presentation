public interface AwesomeVehicle {
    void popASickWheelie();
}

public class GrandmasScooter implements AwesomeVehicle {
    private String modelNumber;
    private int age;

    public GrandmasScooter(String modelNumber, int age) {
        this.modelNumber = modelNumber;
        this.age = age;
    }

    @Override
    public void popASickWheelie() {
        System.out.printf("Why does a %i year old look so fly right now?\n", this.age);
    }
}

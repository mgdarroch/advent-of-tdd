package org.advent.day1;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        ArrayList<Elf> elfList = new ArrayList<>();
        try {
            Scanner scanner = new Scanner(new File("src/main/resources/day1-elf-calories.txt"));

            ArrayList<Integer> intList = new ArrayList<>();
            while (scanner.hasNextLine()) {
                String line = scanner.nextLine();
                if (line.isEmpty()) {
                    Elf elf = new Elf();
                    for (Integer integer : intList) {
                        elf.addCalories(integer);
                    }
                    elfList.add(elf);
                    intList.clear();
                } else {
                    Integer calories = Integer.parseInt(line);
                    intList.add(calories);
                }
            }

            scanner.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }

        for (Elf elf : elfList) {
            System.out.println(elf.getTotalCalories());
        }
    }
}
